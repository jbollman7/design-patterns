//Demonstrates the Decorator Pattern by allowing you to dynamically add behavior (in this case, additional ingredients) 
//to beverage objects without modifying their structure. The use of composition 
//(where decorators hold a reference to a Beverage object) is a key aspect of the pattern,

Beverage decaff = new Decaff();

// Dynamically add behavior by wrapping the Decaff with two Mocha decorators and one Espresso decorator.
decaff = new MochaDecorator(decaff);
decaff = new MochaDecorator(decaff);
decaff = new EspressoDecorator(decaff);

// Output the description and total cost of the decorated beverage.
Console.WriteLine(decaff.getDescription() + $"Total: {decaff.getCost()}");

// Create a DarkRoast beverage with double shot espresso.
Beverage darkRoast = new EspressoDecorator(new EspressoDecorator(new DarkRoast()));

// Output the description and total cost of the double shot espresso dark roast.
Console.WriteLine(darkRoast.getDescription() + $"Total: {darkRoast.getCost()}");

public interface Beverage
{
    public string getDescription();
    public double getCost();
}


public class DarkRoast : Beverage
{

    public string getDescription()
    {
        return "DarkRoast: ";
    }

    public double getCost()
    {
        return 1.00;
    }
}

public class Decaff : Beverage
{

    public string getDescription()
    {
        return "Decaff: ";
    }

    public double getCost()
    {
        return 1.00;
    }
}

public class EspressoDecorator : Beverage
{
    Beverage beverage;
    public EspressoDecorator(Beverage beverage)
    {
        this.beverage = beverage;
    }

    public string getDescription()
    {
        return beverage.getDescription() + "Espresso, ";
    }

    public double getCost()
    {
        return beverage.getCost() + 1.00;
    }
}

public class MochaDecorator : Beverage
{
    Beverage beverage;

    public MochaDecorator(Beverage beverage)
    {
        this.beverage = beverage;
    }

    public string getDescription()
    {
        return beverage.getDescription() + "Mocha, ";
    }

    public double getCost()
    {
        return beverage.getCost() + 2.00;
    }
}
