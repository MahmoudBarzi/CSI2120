import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Comparator;

public class Tree {
	String theFile = "/Users/mahmoudasadzadeh1/Documents/workspace/Wading Pools/src/theData.txt";
	String theLine = null;
	Pool root;
	Pool rootsChild;
	ArrayList<Pool> allPools = new ArrayList<Pool>();
	ArrayList<Pool> thePath; 
	public void readTheFile(){//reads the file and builds the objects 
		try {
			FileReader fileReader = new FileReader(theFile);
			BufferedReader bufferedReader =  new BufferedReader(fileReader);
			while((theLine = bufferedReader.readLine()) != null) { 
				String [] theInfo = theLine.split(",") ;
				String poolName = theInfo[0];
				double latitudes = Double.parseDouble(theInfo[1]);
				double longitudes = Double.parseDouble(theInfo[2]);
				Pool pool = new Pool(poolName,latitudes, longitudes);//creating all my pools
				allPools.add(pool);
				allPools.sort(Comparator.comparingDouble(Pool::getLon));//the array list is sorted from west to east now
			}
			bufferedReader.close();   
			for(int i =0; i <=allPools.size()-1;i++){
				System.out.println("tree("+"'"+allPools.get(i).getName()+"'"+","+"["+"]"+").");//+allPools.get(i).getLat()+","+allPools.get(i).getLon()+"]"+").");
			}
			
		}
		catch(FileNotFoundException ex) {System.out.println( " Can't Open The File '" +theFile + "'");}
		catch(IOException ex) { System.out.println("ERROR, Can't Read The File '"+ theFile + "'");}	
	}	
	public void theTree(){
		root = allPools.get(0);
		int flag = 0;
		for(int i = 1; i <=allPools.size()-1 ; i++  ){
			//Sdtr stands for shortest distance to root
			double tempSdtr = findDistance(root, allPools.get(1));//initializing the tempSdtr to  the distance between root and the next west most pool after the root
			double current =findDistance(root, allPools.get(i)); //current is the distance between root and the next pool in the Array list
			if(current < tempSdtr){
				tempSdtr = current;
				flag = i ;
			}
		}
		rootsChild = allPools.get(flag);
		rootsChild.setHasParent(true);
		root.setChilds(rootsChild);
		rootsChild.setDistanceFromParent(findDistance(root, rootsChild)); //special case: set the distance of the root and its closet child
		for(int i = 1; i<=allPools.size()-1; i++){//pick the next most western pool from allPools 
			double tempDistance = findDistance(allPools.get(i), allPools.get(i-1));
			int flag1= 0;
			for(int j = i-1; j >= 0; j--){//check the distance of i with all the previous pools
				double currentDistance =findDistance(allPools.get(i),allPools.get(j));
				if(currentDistance < tempDistance){
					tempDistance = currentDistance;
					flag1 = j;//found the shortest distance of i with the other pools  
				}
			}
			if(!(allPools.get(i).getHasParent())){//if this pool i does not have a parent then 
				allPools.get(flag1).setChilds(allPools.get(i));//set this pool that we found as its parent
				allPools.get(i).setHasParent(true);//this pool has a parent now.
				allPools.get(i).setDistanceFromParent(findDistance(allPools.get(i), allPools.get(flag1)));//keeping track of the distance between the pool and its parent
			}
		}
	}
	public ArrayList<Pool>  preOrderTraversal(Pool pool){
		ArrayList<Pool> tempPath = new ArrayList<Pool>();
		if(pool == null ) {
			return tempPath;
		}
		tempPath.add(pool);
		for(Pool child : pool.getChilds()){
			tempPath.addAll(preOrderTraversal(child));
		}
		thePath = tempPath;
		return thePath;
	}
	public void showResults(){
		double distanceTraveled = thePath.get(0).getDistanceFromParent();//initializing
		for(int i = 0; i <= thePath.size()-1; i++){
			distanceTraveled += thePath.get(i).getDistanceFromParent();// accumulating
			System.out.println(thePath.get(i).getName() + "    "+distanceTraveled );
		}
	}
	public double findDistance(Pool location1, Pool location2){
		double latitude1 = Math.PI *(location1.getLat()/180);
		double longitude1 = Math.PI *(location1.getLon()/180);
		double latitude2 = Math.PI *(location2.getLat()/180);
		double longitude2 = Math.PI *(location2.getLon()/180);
		double distance = 2 * Math.sin(Math.sqrt(( Math.pow( ( Math.sin ( ( (latitude1-latitude2) / 2)))  , 2) )+
				Math.cos(latitude1)*Math.cos(latitude2)*( Math.pow( ( Math.sin ( ( (longitude1-longitude2) / 2))) ,2 ) ) ) ) ;
		distance = 6371 * distance;
		return distance;
	}
	public static void main(String[] args) {
		Tree t = new Tree();
		t.readTheFile();
		t.theTree();
		t.preOrderTraversal(t.root);
		t.showResults();
	}
}
