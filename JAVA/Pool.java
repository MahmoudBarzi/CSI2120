import java.util.ArrayList;
public class Pool {
	private String name;
	private double lat;
	private double lon;
	private boolean hasParent = false;
	private double distanceFromParent;
	private ArrayList<Pool> childs = new ArrayList<>();// put the child's of this pool in an array list
	Pool(String name, double lat, double lon){
		this.name = name;
		this.lat = lat;
		this.lon = lon;
	}
	public void setHasParent(boolean hasParent) {this.hasParent = hasParent;}
	public boolean getHasParent() {return hasParent;}
	public String getName() {return name;}
	public double getLat() {return lat;}
	public double getLon() {return lon;}
	public ArrayList<Pool> getChilds() {return childs;}
	public void setChilds(Pool child) {childs.add(child);}//takes a child of type pool, and puts it in the array childs
	public String toString(){return "Pool Name : " + name +" .";}
	public double getDistanceFromParent() {return distanceFromParent;}
	public void setDistanceFromParent(double distanceFromParent) {this.distanceFromParent = distanceFromParent;}	
}
